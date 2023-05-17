import * as core from "@actions/core";
import { Install } from "./install";
import { Plateform } from "./plateform";

/*
https://www.chromium.org/getting-involved/download-chromium/
https://chromiumdash.appspot.com/branches
http://cros-omahaproxy.appspot.com/

112 => 1109224 => 1109220
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/1109220/
113 => 1121455 => 1121455
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/1121454/
114 => 1135570 => 1135561
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/1135561/
*/

class Index {

    async run() {
        return new Promise<void>((resolve, reject) => {
            try {

                // Get input parameters
                const version = core.getInput("version");
    
                // Detect plateform
                const plateform = new Plateform();
                plateform.detectPlateform();

                // Identify to download
                let versionToDownload = "";
                if ( version == "112") {
                    versionToDownload = "1109220";
                } else if ( version == "113" ) {
                    versionToDownload = "1121455";
                } else if ( version == "114" ) {
                    versionToDownload = "1135561";
                } else {
                    throw "Invalid chrome version (112 or 113 or 114 version accepted)";
                }
    
                // Install chrome version on plateform
                new Install().install(versionToDownload, plateform);

                resolve();
                
            } catch (err) {
                core.setFailed(String(err));
                reject(err);
            }
        })
    }
}

new Index().run();