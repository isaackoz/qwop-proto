#!/usr/bin/env -S pnpm tsx
/**
 * Generate exports for package.json to organize ts exports by package/service/version
 * Usage: pnpm tsx script.ts --input <inputDir> --output <outputDir>
 */

import { fs } from "zx";
import { join, dirname, basename } from "path";
import { glob } from "glob";

// Parse command line arguments
function parseArgs() {
  const args = process.argv.slice(2);
  const options: { input?: string; output?: string } = {};

  for (let i = 0; i < args.length; i++) {
    if (args[i] === "--input" && i + 1 < args.length) {
      options.input = args[i + 1];
      i++;
    } else if (args[i] === "--output" && i + 1 < args.length) {
      options.output = args[i + 1];
      i++;
    } else if (args[i] === "-h" || args[i] === "--help") {
      console.log(`
Usage: pnpm tsx script.ts --input <inputDir> --output <outputDir>

Options:
  --input   Directory containing the source package.json
  --output  Directory where to output the generated files
  -h, --help  Show this help message
      `);
      process.exit(0);
    }
  }

  if (!options.input || !options.output) {
    console.error("Error: --input and --output options are required.");
    process.exit(1);
  }

  return {
    inputDir: options.input,
    outputDir: options.output,
  };
}

// Get command line arguments
const { inputDir, outputDir } = parseArgs();

console.log("Input directory:", inputDir);
console.log("Output directory:", outputDir);

// Function to generate exports
async function generateExports() {
  // Read the current package.json
  const packageJsonPath = join(inputDir, "package.json");
  console.log(`Reading package.json from: ${packageJsonPath}`);

  try {
    const packageJson = await fs.readJSON(packageJsonPath);

    // Find all .js files (excluding .d.ts and index.js)
    const jsFiles = await glob("**/*.js", {
      cwd: outputDir,
      ignore: ["**/*.d.ts", "**/index.js"],
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
      const indexPath = join(outputDir, dir, "index.js");
      const indexDtsPath = join(outputDir, dir, "index.d.ts");

      // Generate index.js content
      const jsExports = files
        .map((file) => `export * from './${basename(file, ".js")}';`)
        .join("\n");

      // Generate index.d.ts content
      const dtsExports = files
        .map((file) => `export * from './${basename(file, ".js")}';`)
        .join("\n");

      // Write index files
      await fs.writeFile(indexPath, jsExports + "\n");
      await fs.writeFile(indexDtsPath, dtsExports + "\n");

      console.log(`Created index files in ${dir}`);

      // Create the export path
      const exportPath = `./${dir}`;
      const filePath = `./${join(dir, "index.js")}`;
      const typeFilePath = filePath.replace(".js", ".d.ts");

      exports[exportPath] = {
        types: typeFilePath,
        import: filePath,
        require: filePath,
      };
    }

    // Update package.json
    packageJson.exports = exports;

    // Write back to package.json and move it to output directory
    const outputPackageJsonPath = join(outputDir, "package.json");
    await fs.writeJSON(outputPackageJsonPath, packageJson, { spaces: 2 });
    console.log(
      `Updated package.json with exports and saved to: ${outputPackageJsonPath}`
    );
  } catch (err) {
    console.error(`Error: ${err.message}`);
    process.exit(1);
  }
}

generateExports();
