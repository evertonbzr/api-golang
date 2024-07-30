CREATE TABLE IF NOT EXISTS "todos" (
	"id" serial PRIMARY KEY NOT NULL,
	"full_name" varchar(255),
	"description" text,
	"status" varchar(100),
	"user_id" serial NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "todos" ADD CONSTRAINT "todos_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
