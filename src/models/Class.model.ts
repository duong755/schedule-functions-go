import { Document, Schema, model } from "mongoose";

interface IClass extends Document {
  MaMH: string;
  TenMonHoc: string;
  TinChi: number;
  MaLopMH: string;
  GiaoVien: string;
  SoSV: string;
  Buoi: string;
  Thu: string;
  Tiet: string;
  GiangDuong: string;
  GhiChu: string;
}

const ClassSchema = new Schema<IClass>(
  {
    MaMH: String,
    TenMonHoc: String,
    TinChi: Number,
    MaLopMH: String,
    GiaoVien: String,
    SoSV: String,
    Buoi: String,
    Thu: String,
    Tiet: String,
    GiangDuong: String,
    GhiChu: String,
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

export const ClassModel = model<IClass>("class", ClassSchema, "Class");
