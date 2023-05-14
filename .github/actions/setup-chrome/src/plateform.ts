import os from "os";

export class Plateform {
    private system : string;
    private architecture : string;

    constructor () {
        this.architecture = "";
        this.system = "";
    }

    public detectPlateform() {
        this.system = this.detectSystem();
        this.architecture = this.detectArchitecture();
    }

    private detectSystem() {
        const platform = os.platform();
        switch (platform) {
            case "linux":
            return SYSTEM_TYPE.LINUX;
            case "darwin":
            return SYSTEM_TYPE.DARWIN;
            case "win32":
            return SYSTEM_TYPE.WINDOWS;
        }
        throw new Error(`Unsupported platform: ${platform}`);
    }

    private detectArchitecture() {
        const arch = os.arch();
        switch (arch) {
            case "x32":
            return ARCHITECTURE_TYPE.I686;
            case "x64":
            return ARCHITECTURE_TYPE.AMD64;
        }
        throw new Error(`Unsupported arch: ${arch}`);
    }

    public getSystem() : string{
        return this.system;
    }

    public getArchitecture() : string {
        return this.architecture;
    }
}

export const SYSTEM_TYPE = {
    DARWIN: "darwin",
    LINUX: "linux",
    WINDOWS: "windows",
}

export const ARCHITECTURE_TYPE = {
    AMD64: "amd64",
    I686: "i686",
    ARM64: "arm64",
};