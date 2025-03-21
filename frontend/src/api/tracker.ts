import { apiClient } from "./client";

export const sendTracker = async (
  deviceToken: string,
  position: [number, number]
) => {
  console.log("sendTracker deviceToken: ", deviceToken);
  console.log("sendTracker position: ", position);

  const response = await apiClient.post("api/tracker", {
    deviceToken: deviceToken,
    position: position,
  });
};
