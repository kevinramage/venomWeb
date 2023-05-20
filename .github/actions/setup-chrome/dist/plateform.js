"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.ARCHITECTURE_TYPE = exports.SYSTEM_TYPE = exports.Plateform = void 0;
const os_1 = __importDefault(require("os"));
class Plateform {
    constructor() {
        this.architecture = "";
        this.system = "";
    }
    detectPlateform() {
        this.system = this.detectSystem();
        this.architecture = this.detectArchitecture();
    }
    detectSystem() {
        const platform = os_1.default.platform();
        switch (platform) {
            case "linux":
                return exports.SYSTEM_TYPE.LINUX;
            case "darwin":
                return exports.SYSTEM_TYPE.DARWIN;
            case "win32":
                return exports.SYSTEM_TYPE.WINDOWS;
        }
        throw new Error(`Unsupported platform: ${platform}`);
    }
    detectArchitecture() {
        const arch = os_1.default.arch();
        switch (arch) {
            case "x32":
                return exports.ARCHITECTURE_TYPE.I686;
            case "x64":
                return exports.ARCHITECTURE_TYPE.AMD64;
        }
        throw new Error(`Unsupported arch: ${arch}`);
    }
    getSystem() {
        return this.system;
    }
    getArchitecture() {
        return this.architecture;
    }
}
exports.Plateform = Plateform;
exports.SYSTEM_TYPE = {
    DARWIN: "darwin",
    LINUX: "linux",
    WINDOWS: "windows",
};
exports.ARCHITECTURE_TYPE = {
    AMD64: "amd64",
    I686: "i686",
    ARM64: "arm64",
};
