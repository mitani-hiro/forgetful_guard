import { create } from "zustand";
import { sendTracker } from "../api/tracker";

interface TrackerState {
  sendTracker: (
    deviceToken: string,
    position: [number, number]
  ) => Promise<void>;
}

export const useTrackerStore = create<TrackerState>(() => ({
  position: [],

  sendTracker: async (deviceToken, position) => {
    try {
      await sendTracker(deviceToken, position);
    } catch (error) {
      alert(error);
    }
  },
}));
