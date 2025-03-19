import { create } from "zustand";
import { sendTracker } from "../api/tracker";

interface TrackerState {
  sendTracker: (position: [number, number]) => Promise<void>;
}

export const useTrackerStore = create<TrackerState>(() => ({
  position: [],

  sendTracker: async (position) => {
    try {
      await sendTracker(position);
    } catch (error) {
      alert(error);
    }
  },
}));
