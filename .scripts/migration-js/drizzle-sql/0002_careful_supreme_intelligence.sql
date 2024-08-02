ALTER TABLE "todos" ADD COLUMN "title" varchar(255);--> statement-breakpoint
ALTER TABLE "todos" DROP COLUMN IF EXISTS "full_name";