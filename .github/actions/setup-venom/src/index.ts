import * as core from "@actions/core";
import { Install } from "./install";

class Index {

    async run() {
        return new Promise<void>((resolve, reject) => {
            try {

                // Get input parameters
                const venomVersion = core.getInput("venomVersion");
                const venomWebVersion = core.getInput("venomWebVersion");
                const targetPlateform = core.getInput("plateform");

                // Install chrome version on plateform
                new Install().install(venomVersion, venomWebVersion, targetPlateform);

                resolve();
                
            } catch (err) {
                core.setFailed(String(err));
                reject(err);
            }
        })
    }
}

new Index().run();