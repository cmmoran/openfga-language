import * as path from "path";
import * as fs from "fs";
import * as yaml from "js-yaml";

interface ValidTestCase {
  name: string;
  dsl: string;
  json: string;
  skip: boolean;
}

export function loadValidTransformerTestCases(): ValidTestCase[] {
  const testDataPath = path.join(__dirname, "../../../tests", "data", "transformer");
  const entries = fs.readdirSync(testDataPath, { withFileTypes: true });

  const testCases: ValidTestCase[] = [];

  for (const entry of entries) {
    if (!entry.isDirectory()) {
      continue;
    }

    const testCase: Partial<ValidTestCase> = {
      name: entry.name,
    };

    try {
      const skipFile = fs.readFileSync(path.join(testDataPath, testCase.name!, "test.skip"));
      if (skipFile) {
        testCase.skip = true;
      }
    } catch (e) {
      // do nothing
    }

    const jsonData = fs.readFileSync(path.join(testDataPath, testCase.name!, "authorization-model.json"));
    testCase.json = jsonData.toString("utf8");

    const dslData = fs.readFileSync(path.join(testDataPath, testCase.name!, "authorization-model.fga"));
    testCase.dsl = dslData.toString("utf8");

    testCases.push(testCase as ValidTestCase);
  }

  return testCases;
}

interface InvalidDslSyntaxTestCase {
  name: string;
  dsl: string;
  valid: boolean;
  error_message: string;
}

export function loadInvalidDslSyntaxTestCases(): InvalidDslSyntaxTestCase[] {
  const jsonData = fs.readFileSync(path.join(__dirname, "../../../tests", "data", "dsl-syntax-validation.json"));

  return JSON.parse(jsonData.toString("utf8")) as InvalidDslSyntaxTestCase[];
}

interface MultipleInvalidDslSyntaxTestCase extends InvalidDslSyntaxTestCase {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  expectedError: any;
}

export function loadDslSyntaxErrorTestCases(): MultipleInvalidDslSyntaxTestCase[] {
  const yamlData = yaml.loadAll(fs.readFileSync(
    path.join(__dirname, "../../../tests", "data", "dsl-syntax-validation-cases.yaml"), "utf-8"));
  return yamlData as MultipleInvalidDslSyntaxTestCase[];
}