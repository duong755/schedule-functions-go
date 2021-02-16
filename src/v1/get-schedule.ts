import { NowRequest, NowResponse } from "@vercel/node";

import { connectToMongoAtlas } from "../common/connector";

const getSchedule = async (_req: NowRequest, res: NowResponse): Promise<void> => {
  await connectToMongoAtlas();
  // const { studentCode } = req.body as { studentCode: string };
  res.json({});
};

export default getSchedule;
