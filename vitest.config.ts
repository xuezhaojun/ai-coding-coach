import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    include: [
      "templates/problems/**/typescript/*.test.ts",
      "my-progress/**/typescript/*.test.ts",
    ],
    globals: true,
  },
});
