#!/usr/bin/env -S pnpm tsx
/**
 * Generate exports for package.json to organize ts exports by package/service/version
 */

import { fs, $ } from "zx";
import { resolve, join, dirname, basename } from "path";
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

  // Find all .js files (excluding .d.ts and index.js)
  const jsFiles = await glob("**/*.js", { 
    cwd: distDir,
    ignore: ["**/*.d.ts", "**/index.js"]
  });

  // Group files by directory
  const dirMap = new Map<string, string[]>();
  
  jsFiles.forEach((file) => {
    const dir = dirname(file);
    if (!dirMap.has(dir)) {
      dirMap.set(dir, []);
    }
    dirMap.get(dir)!.push(file);
  });

  // Generate exports object
  const exports = {};

  // Generate index files for each directory
  for (const [dir, files] of dirMap) {
    const indexPath = join(distDir, dir, 'index.js');
    const indexDtsPath = join(distDir, dir, 'index.d.ts');
    
    // Generate index.js content
    const jsExports = files
      .map(file => `export * from './${basename(file, '.js')}';`)
      .join('\n');
    
    // Generate index.d.ts content
    const dtsExports = files
      .map(file => `export * from './${basename(file, '.js')}';`)
      .join('\n');
    
    // Write index files
    await fs.writeFile(indexPath, jsExports + '\n');
    await fs.writeFile(indexDtsPath, dtsExports + '\n');
    
    console.log(`Created index files in ${dir}`);

    // Create the export path
    const exportPath = `./${dir}`;
    const filePath = `./${join(dir, 'index.js')}`;
    const typeFilePath = filePath.replace(".js", ".d.ts");

    exports[exportPath] = {
      types: typeFilePath,
      import: filePath,
      require: filePath,
    };
  }

  // Update package.json
  packageJson.exports = exports;

  // Write back to package.json and move it to dist
  await fs.writeJSON("dist/package.json", packageJson, { spaces: 2 });
  console.log("Updated package.json with exports for:", exports);
}

await generateExports();
