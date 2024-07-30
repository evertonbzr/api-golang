import { drizzle } from "drizzle-orm/postgres-js";
import { migrate } from "drizzle-orm/postgres-js/migrator";
import postgres from "postgres";

const sql = postgres(String(process.env.DATABASE_URL), { max: 1 });

const db = drizzle(sql);

async function main() {
  await migrate(db, { migrationsFolder: "drizzle-sql" });

  await sql.end();
}

main();
