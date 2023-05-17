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
            await exec.exec("go", ["get", `kevinramage/venomWeb@${venomWebVersion}`], options);

            // Checkout venom web
            //await exec.exec("git", ["clone", "-b", venomWebVersion, "https://github.com/kevinramage/venomWeb.git", "venomWeb"]);

            // Compile venom for target plateform (windows-latest, ubuntu-latest, macos-latest)
            core.info(`Compile`);
            options.cwd = "./venom/cmd/venom";
            let targetPlateform = "";
            if (plateform == "ubuntu-latest") {
                targetPlateform = "OS=linux";
            } else if (plateform == "macos-latest") {
                targetPlateform = "OS=darwin"
            } else if (plateform == "windows-latest") {
                targetPlateform = "OS=windows";
            }
            await exec.exec("make", ["build", targetPlateform, "ARCH=amd64"], options);

            // Rename build
            await exec.exec("mv", ["dist/venom.*", "venom"], options);

            resolve();

            } catch (err) {
                core.error(`An error during installation of venom: ${err}`)
                reject(err);
            }
        });
    }
}