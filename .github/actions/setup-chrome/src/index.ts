import * as core from "@actions/core";
import * as exec from "@actions/exec";
import { Install } from "./install";
import { Plateform } from "./plateform";

class Index {

    async run() {
        return new Promise<void>((resolve, reject) => {
            try {

                // Get input parameters
                const version = core.getInput("version") || "latest";
    
                // Detect plateform
                const plateform = new Plateform();
                plateform.detectPlateform();
    
                // Install chrome version on plateform
                new Install().install("1125695", plateform);

                resolve();
                
            } catch (err) {
                core.setFailed(String(err));
                reject(err);
            }
        })
    }
}

new Index().run();