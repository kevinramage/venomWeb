import * as core from "@actions/core";
import * as exec from "@actions/exec";

export class Install {
    public install(venomVersion : string, venomWebVersion: string, plateform: string) {
        return new Promise<void>(async(resolve, reject) => {
            core.info(`Install venom version ${venomVersion} with venomWeb version ${venomWebVersion}`);
            try {

            // Checkout venom
            core.info(`Clone venom repository`);
            await exec.exec("git", ["clone", "-b", venomVersion, "https://github.com/ovh/venom.git", "venom"]);

            // Get venom web specific version
            core.info(`Get venomWeb dependency`);
            let options : any = {};
            options.cwd = "./venom";
            const dependency = "github.com/kevinramage/venomWeb@" + venomWebVersion;
            core.info(dependency);
            await exec.exec("go", ["get", dependency], options);

            // Checkout venom web
            //await exec.exec("git", ["clone", "-b", venomWebVersion, "https://github.com/kevinramage/venomWeb.git", "venomWeb"]);

            // Compile venom for target plateform (windows-latest, ubuntu-latest, macos-latest)
            core.info(`Compile`);
            options.cwd = "./venom/cmd/venom";
            let targetPlateform = "", execName = "";
            if (plateform == "ubuntu-latest") {
                targetPlateform = "OS=linux";
                execName = "venom.linux-amd64";
            } else if (plateform == "macos-latest") {
                targetPlateform = "OS=darwin"
                execName = "venom.darwin-amd64"
            } else if (plateform == "windows-latest") {
                targetPlateform = "OS=windows";
                execName = "venom.windows-amd64"
            }
            await exec.exec("make", ["build", targetPlateform, "ARCH=amd64"], options);

            // Rename build
            await exec.exec("mv", ["dist/" + execName , "venom"], options);

            resolve();

            } catch (err) {
                core.error(`An error during installation of venom: ${err}`)
                reject(err);
            }
        });
    }
}