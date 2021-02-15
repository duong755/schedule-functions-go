import { Document, model, Schema } from "mongoose";

interface ISchedule extends Document {
  NgaySinh: string;
  LopKhoaHoc: string;
  MaLMH: string;
  TenMonHoc: string;
  Nhom: string;
  SoTinChi: string;
  GhiChu: string;
  MaSV: string;
  HoVaTen: string;
}

const ScheduleSchema = new Schema<ISchedule>(
  {
    NgaySinh: String,
    LopKhoaHoc: String,
    MaLMH: String,
    Nhom: String,
    SoTinChi: String,
    GhiChu: String,
    MaSV: String,
    HoVaTen: String,
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

export const ScheduleModel = model<Document<ISchedule>>("schedule", ScheduleSchema, "Schedule");
