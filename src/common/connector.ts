import { config } from "dotenv";
import { connect } from "mongoose";

config();

export async function connectToMongoAtlas(): Promise<void> {
  connect(process.env.DATABASE_CONNECTION_STRING as string, { useNewUrlParser: true });
}
