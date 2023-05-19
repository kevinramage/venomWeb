import * as core from "@actions/core";
import { Install } from "./install";
import { Plateform, SYSTEM_TYPE } from "./plateform";

/*
https://www.chromium.org/getting-involved/download-chromium/
https://chromiumdash.appspot.com/branches
http://cros-omahaproxy.appspot.com/

112 => 1109224 => 1109220 (Unix), 1109213 (Mac), 1109208 (Win)
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/1109220/
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Mac/1109213/
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Win_x64/1109208/
113 => 1121455 => 1121454 (Unix), 1121448 (Mac), 1121434 (Win)
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/1121454/
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Mac/1121448/
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Win_x64/1121434/
114 => 1135570 => 1135561 (Unix), 1135562 (Mac), 1135559 (Win)
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Linux_x64/1135561/
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Mac/1135562/
https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html?prefix=Win_x64/1135559/
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
                    if (plateform.getSystem() == SYSTEM_TYPE.DARWIN) {
                        versionToDownload = "1109213";
                    } else if (plateform.getSystem() == SYSTEM_TYPE.LINUX) {
                        versionToDownload = "1109220";
                    } else if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                        versionToDownload = "1109208";
                    }
                } else if ( version == "113" ) {
                    if (plateform.getSystem() == SYSTEM_TYPE.DARWIN) {
                        versionToDownload = "1121448";
                    } else if (plateform.getSystem() == SYSTEM_TYPE.LINUX) {
                        versionToDownload = "1121454";
                    } else if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                        versionToDownload = "1121434";
                    }
                } else if ( version == "114" ) {
                    if (plateform.getSystem() == SYSTEM_TYPE.DARWIN) {
                        versionToDownload = "1135562";
                    } else if (plateform.getSystem() == SYSTEM_TYPE.LINUX) {
                        versionToDownload = "1135561";
                    } else if (plateform.getSystem() == SYSTEM_TYPE.WINDOWS) {
                        versionToDownload = "1135559";
                    }
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