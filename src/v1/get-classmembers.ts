import { NowRequest, NowResponse } from "@vercel/node";

const getClassMembers = (_: NowRequest, res: NowResponse): void => {
  res.json({});
};

export default getClassMembers;
