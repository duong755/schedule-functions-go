import { NowRequest, NowResponse } from "@vercel/node";

const getSchedule = (req: NowRequest, res: NowResponse) => {
  console.log(req.query);
  res.json({
    message: "Hello World"
  });
};

export default getSchedule;
