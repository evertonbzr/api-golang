import { relations } from "drizzle-orm";
import { pgTable, serial, text, timestamp, varchar } from "drizzle-orm/pg-core";
import { users } from "./users";

export const todos = pgTable("todos", {
  id: serial("id").primaryKey(),
  title: varchar("full_name", { length: 255 }),
  description: text("description"),
  status: varchar("status", { length: 100 }),
  userId: serial("user_id").references(() => users.id, {
    onDelete: "set null",
  }),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at").defaultNow(),
});

export const todosRelations = relations(todos, ({ one }) => ({
  user: one(users, {
    fields: [todos.userId],
    references: [users.id],
    relationName: "todos_user",
  }),
}));
