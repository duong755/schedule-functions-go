import { NowRequest, NowResponse } from "@vercel/node";

const getSchedule = (_: NowRequest, res: NowResponse): void => {
  res.json({
    message: "Hello World",
  });
};

export default getSchedule;
