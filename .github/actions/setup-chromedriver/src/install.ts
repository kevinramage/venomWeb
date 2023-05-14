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
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Win%2F${version}%2Fchromedriver_win32.zip?alt=media`;
            } else if (plateform.getArchitecture() == ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Win_x64%2F${version}%2Fchromedriver_win32.zip?alt=media`;
            }
        } else if (plateform.getSystem() == SYSTEM_TYPE.LINUX) {
            if (plateform.getArchitecture() == ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux%2F${version}%2Fchromedriver-linux.zip?alt=media`;
            } else if (plateform.getArchitecture() == ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux_x64%2F${version}%2Fchromedriver_linux64.zip?alt=media`;
            }
        } else if (plateform.getSystem() == SYSTEM_TYPE.DARWIN) {
            if (plateform.getArchitecture() == ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Mac%2F${version}%2Fchromedriver_mac.zip?alt=media`;
            } else if (plateform.getArchitecture() == ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Mac%2F${version}%2Fchromedriver_mac64.zip?alt=media`;
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
                await exec.exec("unzip", ["-d", "/opt/chromedriver", "-j", archivePath])

                // Remove archive
                await fs.promises.unlink(archivePath);

                // Rename folder
                await fs.promises.rename("/opt/chromedriver/chromedriver_linux64", "/opt/chromedriver/chromedriver");

                // Add chrome to path
                core.info(`Add chrome binary to path`);
                //await exec.exec("ln", ["-s", "/opt/chrome/chrome-linux/chrome", "chrome"]);
                core.addPath("/opt/chromedriver/chromedriver")
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
                await exec.exec("unzip", ["-d", "/opt/chromedriver", archivePath])

                // Remove archive
                await fs.promises.unlink(archivePath);

                // Rename folder
                await fs.promises.rename("/opt/chromedriver/chromedriver_mac64", "/opt/chromedriver/chromedriver");

                // Add chrome to path
                core.info(`Add chrome binary to path`);
                //await exec.exec("ln", ["-s", "/opt/chrome/chrome-linux/chrome", "chrome"]);
                core.addPath("/opt/chromedriver/chromedriver")
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
                await fs.promises.rename("/opt/chromedriver/chromedriver_win32", "/opt/chromedriver/chromedriver");

                // Add chrome to path
                core.info(`Add chrome binary to path`);
                //await exec.exec("ln", ["-s", "/opt/chrome/chrome-linux/chrome", "chrome"]);
                core.addPath(`{destination}\\chromedriver`)
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
                } else if (plateform.getArchitecture() == SYSTEM_TYPE.DARWIN) {
                    await this.installDarwin(archivePath);
                // Install binary (Windows)
                } else if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                    await this.installWindows(archivePath, plateform);
                }

                resolve();
            } catch (err) {
                core.error(String(err));
                reject(err);
            }
        });
    }
}