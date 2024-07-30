import { defineConfig } from "drizzle-kit";

export default defineConfig({
  schema: "./src/db/schema/index.ts",
  out: "./drizzle-sql",
  dialect: "postgresql",
  dbCredentials: {
    url: String(process.env.DATABASE_URL),
  },
});
