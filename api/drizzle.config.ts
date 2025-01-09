import "dotenv/config";
import { Config, defineConfig } from "drizzle-kit";

const config: Config = {
  dialect: "sqlite",
  out: "./drizzle",
  schema: "./src/db/schema.ts",
  driver: "d1-http",
  dbCredentials: {
    accountId: process.env.CLOUDFLARE_ACCOUNT_ID!,
    databaseId: process.env.CLOUDFLARE_DATABASE_ID!,
    token: process.env.CLOUDFLARE_D1_TOKEN!,
  },
};

export default defineConfig(config);
