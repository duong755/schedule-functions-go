import { Document, model, Schema } from "mongoose";

interface ITrack extends Document {
  studentCode: string;
  accessAt: Date;
  ipAddress: string;
}

const TrackSchema = new Schema<ITrack>(
  {
    studentCode: String,
    accessAt: Date,
    ipAddress: String,
  },
  {
    toJSON: {
      virtuals: true,
      versionKey: false,
      transform: (_doc, ret) => {
        ret.id = ret._id;
        delete ret._id;
      },
    },
  }
);

export const TrackModel = model<ITrack>("track", TrackSchema, "Track");
