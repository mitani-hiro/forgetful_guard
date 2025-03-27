import { apiClient } from "./client";
import { components } from "./openapi.gen";

export const sendTracker = async (
  deviceToken: string,
  position: [number, number]
) => {
  console.log("sendTracker deviceToken: ", deviceToken);
  console.log("sendTracker position: ", position);

  const tracker: components["schemas"]["Tracker"] = {
    deviceToken: deviceToken,
    position: position,
  };

  const response = await apiClient.post("/api/tracker", tracker);
};
