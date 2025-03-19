#!/usr/bin/env -S pnpm tsx
/**
 * Generate exports for package.json to organize ts exports by package/service/version
 */

import { fs, $ } from "zx";
import { resolve, join, dirname } from "path";
import { glob } from "glob";
import { fileURLToPath } from "url";

const __dirname = dirname(fileURLToPath(import.meta.url));

// Define the directory to scan
const distDir = resolve(__dirname, "../dist");
console.log("dist dir", distDir);

// Function to generate exports
async function generateExports() {
  // Read the current package.json
  const packageJson = await fs.readJSON("package.json", { dirname });

  // Find all pb.js files
  const pbFiles = await glob("**/*_pb.js", { cwd: distDir });

  // Generate exports object
  const exports = {};

  pbFiles.forEach((file) => {
    // Get the directory path without the file name
    const dirPath = dirname(file);

    // Create the export path (without _pb suffix)
    const exportPath = `./${dirPath}`;

    // Create the file path
    const filePath = `./${file}`;
    console.log("fp is", filePath);
    const typeFilePath = filePath.replace(".js", ".d.ts");

    exports[exportPath] = {
      types: typeFilePath,
      import: filePath,
      require: filePath,
    };
  });

  // Update package.json
  packageJson.exports = exports;

  // Write back to package.json and move it to dist
  await fs.writeJSON("dist/package.json", packageJson, { spaces: 2 });
  console.log("Updated package.json with exports for:", exports);
}

await generateExports();
