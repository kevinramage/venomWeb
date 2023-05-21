"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Install = void 0;
const tc = __importStar(require("@actions/tool-cache"));
const exec = __importStar(require("@actions/exec"));
const core = __importStar(require("@actions/core"));
const fs_1 = __importDefault(require("fs"));
const plateform_1 = require("./plateform");
class Install {
    getDownloadUrl(version, plateform) {
        if (plateform.getSystem() == plateform_1.SYSTEM_TYPE.WINDOWS) {
            if (plateform.getArchitecture() == plateform_1.ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Win%2F${version}%2Fchrome-win.zip?alt=media`;
            }
            else if (plateform.getArchitecture() == plateform_1.ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Win_x64%2F${version}%2Fchrome-win.zip?alt=media`;
            }
        }
        else if (plateform.getSystem() == plateform_1.SYSTEM_TYPE.LINUX) {
            if (plateform.getArchitecture() == plateform_1.ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux%2F${version}%2Fchrome-linux.zip?alt=media`;
            }
            else if (plateform.getArchitecture() == plateform_1.ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux_x64%2F${version}%2Fchrome-linux.zip?alt=media`;
            }
        }
        else if (plateform.getSystem() == plateform_1.SYSTEM_TYPE.DARWIN) {
            if (plateform.getArchitecture() == plateform_1.ARCHITECTURE_TYPE.I686) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Mac%2F${version}%2Fchrome-mac.zip?alt=media`;
            }
            else if (plateform.getArchitecture() == plateform_1.ARCHITECTURE_TYPE.AMD64) {
                return `https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Mac%2F${version}%2Fchrome-mac.zip?alt=media`;
            }
        }
        throw `Unsupported plateform: ${plateform.getSystem()} - ${plateform.getArchitecture()}`;
    }
    downloadSetup(version, plateform) {
        return new Promise((resolve, reject) => __awaiter(this, void 0, void 0, function* () {
            core.info(`Download setup for version ${version} and plateform ${plateform.getSystem()} - ${plateform.getArchitecture()}`);
            const url = this.getDownloadUrl(version, plateform);
            try {
                const archivePath = yield tc.downloadTool(url);
                resolve(archivePath);
            }
            catch (err) {
                core.error(`An error during download setup: ${err}`);
                reject(err);
            }
        }));
    }
    installUnix(archivePath) {
        return new Promise((resolve, reject) => __awaiter(this, void 0, void 0, function* () {
            core.info(`Install to unix system: ${archivePath}`);
            try {
                // Unarchive
                yield exec.exec("unzip", ["-d", "/opt/chrome", archivePath]);
                // Remove archive
                yield fs_1.default.promises.unlink(archivePath);
                // Rename folder
                yield fs_1.default.promises.rename("/opt/chrome/chrome-linux", "/opt/chrome/chrome");
                // Add chrome to path
                core.info(`Add chrome binary to path`);
                core.addPath("/opt/chrome/chrome");
                // Display chrome version
                let output = "";
                let options = {};
                options.listeners = {
                    stdout: (data) => {
                        output += data.toString();
                    },
                };
                yield exec.exec("/opt/chrome/chrome/chrome", ["--version"], options);
                core.info("Chrome version: ");
                core.info(output);
                resolve();
            }
            catch (err) {
                core.error(`An error occured during installation: ${err}`);
                reject(err);
            }
        }));
    }
    installDarwin(archivePath) {
        return new Promise((resolve, reject) => __awaiter(this, void 0, void 0, function* () {
            core.info(`Install to darwin system: ${archivePath}`);
            try {
                // Unarchive
                yield exec.exec("sudo unzip", ["-d", "/opt/chrome", archivePath]);
                // Remove archive
                yield fs_1.default.promises.unlink(archivePath);
                // Rename folder
                yield exec.exec("sudo chmod 777 /opt/chrome");
                yield exec.exec("sudo chmod 777 /opt/chrome/chrome-mac");
                yield exec.exec("sudo mv /opt/chrome/chrome-mac /opt/chrome/chrome");
                // Links
                yield exec.exec("sudo chmod 777 /opt/chrome/chrome/Chromium.app/Contents/MacOS/Chromium");
                yield exec.exec("sudo ln -s /opt/chrome/chrome/Chromium.app/Contents/MacOS/Chromium /opt/chrome/chrome/chrome");
                // Add chrome to path
                core.info(`Add chrome binary to path`);
                core.addPath("/opt/chrome/chrome");
                // Display chrome version
                let output = "";
                let options = {};
                options.listeners = {
                    stdout: (data) => {
                        output += data.toString();
                    },
                };
                yield exec.exec("sudo /opt/chrome/chrome/Chromium.app/Contents/MacOS/Chromium", ["--version"], options);
                core.info("Chrome version: ");
                core.info(output);
                resolve();
            }
            catch (err) {
                core.error(`An error occured during installation: ${err}`);
                reject(err);
            }
        }));
    }
    installWindows(archivePath, plateform) {
        return new Promise((resolve, reject) => __awaiter(this, void 0, void 0, function* () {
            core.info(`Install to windows system: ${archivePath}`);
            try {
                // Unarchive
                core.info("Unarchive");
                const destination = plateform.getArchitecture() == plateform_1.ARCHITECTURE_TYPE.AMD64 ? "C:\\Program Files" : "C:\\Program Files (x86)";
                yield exec.exec("7z", ["x", archivePath, `-o${destination}`]);
                // Remove archive
                core.info("Remove archive");
                yield fs_1.default.promises.unlink(archivePath);
                // Rename folder
                core.info("Rename folder");
                yield fs_1.default.promises.rename(destination + "\\chrome-win", destination + "\\chrome");
                // Add chrome to path
                core.info(`Add chrome binary to path`);
                core.addPath("\"C:\\Program Files\\chrome\"");
                yield exec.exec("chrome.exe");
                // Display chrome version
                core.info("Display chrome version");
                let output = "";
                let options = {};
                options.listeners = {
                    stdout: (data) => {
                        output += data.toString();
                    },
                };
                yield exec.exec("powershell (Get-Item C:\\Program` Files\\chrome\\chrome.exe).VersionInfo", [], options);
                resolve();
            }
            catch (err) {
                core.error(`An error occured during installation: ${err}`);
                reject(err);
            }
        }));
    }
    install(version, plateform) {
        return new Promise((resolve, reject) => __awaiter(this, void 0, void 0, function* () {
            core.info(`Install version ${version} for plateform ${plateform.getSystem()} - ${plateform.getArchitecture()}`);
            try {
                // Download version
                const archivePath = yield this.downloadSetup(version, plateform);
                // Install binary (Unix)
                if (plateform.getSystem() == plateform_1.SYSTEM_TYPE.LINUX) {
                    yield this.installUnix(archivePath);
                    // Install binary (Mac)
                }
                else if (plateform.getSystem() == plateform_1.SYSTEM_TYPE.DARWIN) {
                    yield this.installDarwin(archivePath);
                    // Install binary (Windows)
                }
                else if (plateform.getSystem() == plateform_1.SYSTEM_TYPE.WINDOWS) {
                    yield this.installWindows(archivePath, plateform);
                }
                else {
                    throw "Invalid system: " + plateform.getSystem();
                }
                resolve();
            }
            catch (err) {
                core.error(String(err));
                reject(err);
            }
        }));
    }
}
exports.Install = Install;
