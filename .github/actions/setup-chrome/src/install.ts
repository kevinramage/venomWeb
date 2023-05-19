import * as tc from "@actions/tool-cache";
import * as exec from "@actions/exec";
import * as core from "@actions/core";
import os, { version } from "os";
import fs from "fs";
import { ARCHITECTURE_TYPE, Plateform, SYSTEM_TYPE } from "./plateform";

export class Install {

    private getDownloadUrl(version: string, plateform: Plateform) : string {
        if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
            if (plateform.getArchitecture() == ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Win%2F${version}%2Fchrome-win.zip?alt=media`;
            } else if (plateform.getArchitecture() == ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Win_x64%2F${version}%2Fchrome-win.zip?alt=media`;
            }
        } else if (plateform.getSystem() == SYSTEM_TYPE.LINUX) {
            if (plateform.getArchitecture() == ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux%2F${version}%2Fchrome-linux.zip?alt=media`;
            } else if (plateform.getArchitecture() == ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux_x64%2F${version}%2Fchrome-linux.zip?alt=media`;
            }
        } else if (plateform.getSystem() == SYSTEM_TYPE.DARWIN) {
            if (plateform.getArchitecture() == ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Mac%2F${version}%2Fchrome-mac.zip?alt=media`;
            } else if (plateform.getArchitecture() == ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Mac%2F${version}%2Fchrome-mac.zip?alt=media`;
            }
        }
        throw `Unsupported plateform: ${plateform.getSystem()} - ${plateform.getArchitecture()}`;
    }

    private downloadSetup(version: string, plateform : Plateform) {
        return new Promise<string>(async(resolve, reject) => {
            core.info(`Download setup for version ${version} and plateform ${plateform.getSystem()} - ${plateform.getArchitecture()}`)
            const url = this.getDownloadUrl(version, plateform);
            try {
                const archivePath = await tc.downloadTool(url);
                resolve(archivePath);
            } catch (err) {
                core.error(`An error during download setup: ${err}`)
                reject(err);
            }
        })
    }

    private installUnix(archivePath: string) {
        return new Promise<void>(async (resolve, reject) => {
            core.info(`Install to unix system: ${archivePath}`);

            try {
                // Unarchive
                await exec.exec("unzip", ["-d", "/opt/chrome", archivePath])

                // Remove archive
                await fs.promises.unlink(archivePath);

                // Rename folder
                await fs.promises.rename("/opt/chrome/chrome-linux", "/opt/chrome/chrome");

                // Add chrome to path
                core.info(`Add chrome binary to path`);
                core.addPath("/opt/chrome/chrome")

                // Display chrome version
                let output = "";
                let options : any = {};
                options.listeners = {
                    stdout: (data: Buffer) => {
                        output += data.toString();
                    },
                };
                await exec.exec("/opt/chrome/chrome/chrome", ["--version"], options);
                core.info("Chrome version: ");
                core.info(output);

                resolve();

            } catch (err) {
                core.error(`An error occured during installation: ${err}`);
                reject(err);
            }
        });
    }

    private installDarwin(archivePath: string) {
        return new Promise<void>(async (resolve, reject) => {
            core.info(`Install to darwin system: ${archivePath}`);

            try {
                // Unarchive
                await exec.exec("sudo unzip", ["-d", "/opt/chrome", archivePath])

                // Remove archive
                await fs.promises.unlink(archivePath);

                // Rename folder
                await exec.exec("sudo chmod 777 /opt/chrome");
                await exec.exec("sudo chmod 777 /opt/chrome/chrome-mac");
                await exec.exec("sudo mv /opt/chrome/chrome-mac /opt/chrome/chrome");

                // Links
                await exec.exec("sudo chmod 777 /opt/chrome/chrome/Chromium.app/Contents/MacOS/Chromium");
                await exec.exec("sudo ln -s /opt/chrome/chrome/Chromium.app/Contents/MacOS/Chromium /opt/chrome/chrome/chrome");
                
                // Add chrome to path
                core.info(`Add chrome binary to path`);
                core.addPath("/opt/chrome/chrome");

                // Display chrome version
                let output = "";
                let options : any = {};
                options.listeners = {
                    stdout: (data: Buffer) => {
                        output += data.toString();
                    },
                };
                await exec.exec("sudo /opt/chrome/chrome/Chromium.app/Contents/MacOS/Chromium", ["--version"], options);
                core.info("Chrome version: ");
                core.info(output);

                resolve();

            } catch (err) {
                core.error(`An error occured during installation: ${err}`);
                reject(err);
            }
        });
    }

    private installWindows(archivePath: string, plateform: Plateform) {
        return new Promise<void>(async (resolve, reject) => {
            core.info(`Install to windows system: ${archivePath}`);

            try {
                // Unarchive
                const destination = plateform.getArchitecture() == ARCHITECTURE_TYPE.AMD64 ? "C:\\Program Files" : "C:\\Program Files (x86)"
                await exec.exec("7z", ["x", archivePath, `-o${destination}`])

                // Remove archive
                await fs.promises.unlink(archivePath);

                // Rename folder
                await fs.promises.rename( destination + "\\chrome-win", destination + "\\chrome");
                await exec.exec("dir \"C:\\Program Files\"");
                await exec.exec("dir \"C:\\Program Files\\chrome\"");

                // Add chrome to path
                core.info(`Add chrome binary to path`);
                //core.addPath(destination + "\\chrome");

                // Display chrome version
                let output = "";
                let options : any = {};
                options.listeners = {
                    stdout: (data: Buffer) => {
                        output += data.toString();
                    },
                };
                const cmdLine = "cd \"" + destination + "\\chrome\" && powershell (Get-Item chrome.exe).VersionInfo";
                core.info(cmdLine);
                await exec.exec(cmdLine, [], options);
                core.info("Chrome version: ");
                core.info(output);

                resolve();

            } catch (err) {
                core.error(`An error occured during installation: ${err}`);
                reject(err);
            }
        });
    }

    public install(version: string, plateform: Plateform) {
        return new Promise<void>(async (resolve, reject) => {
            core.info(`Install version ${version} for plateform ${plateform.getSystem()} - ${plateform.getArchitecture()}`)

            try {
                // Download version
                const archivePath = await this.downloadSetup(version, plateform);

                // Install binary (Unix)
                if (plateform.getSystem() == SYSTEM_TYPE.LINUX) {
                    await this.installUnix(archivePath);
                // Install binary (Mac)
                } else if (plateform.getSystem() == SYSTEM_TYPE.DARWIN) {
                    await this.installDarwin(archivePath);
                // Install binary (Windows)
                } else if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                    await this.installWindows(archivePath, plateform);
                } else {
                    throw "Invalid system: " + plateform.getSystem();
                }

                resolve();
            } catch (err) {
                core.error(String(err));
                reject(err);
            }
        });
    }
}