import { apiClient } from "./client";

export const sendTracker = async (position: [number, number]) => {
  console.log("sendTracker position: ", position);

  const response = await apiClient.post("api/tracker", {
    deviceID: "device-id-1",
    position: position,
  });
};
