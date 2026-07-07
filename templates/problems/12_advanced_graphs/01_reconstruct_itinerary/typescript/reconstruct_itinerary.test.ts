import { describe, it, expect } from "vitest";
import { findItinerary } from "./reconstruct_itinerary";

describe("findItinerary", () => {
    it.each([
        {
            name: "standard itinerary",
            tickets: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]],
            want: ["JFK", "MUC", "LHR", "SFO", "SJC"],
        },
        {
            name: "lexical order matters",
            tickets: [["JFK", "SFO"], ["JFK", "ATL"], ["SFO", "ATL"], ["ATL", "JFK"], ["ATL", "SFO"]],
            want: ["JFK", "ATL", "JFK", "SFO", "ATL", "SFO"],
        },
        { name: "single ticket", tickets: [["JFK", "ATL"]], want: ["JFK", "ATL"] },
        { name: "round trip", tickets: [["JFK", "ATL"], ["ATL", "JFK"]], want: ["JFK", "ATL", "JFK"] },
        {
            name: "three cities chain",
            tickets: [["JFK", "KUL"], ["JFK", "NRT"], ["NRT", "JFK"]],
            want: ["JFK", "NRT", "JFK", "KUL"],
        },
        {
            name: "duplicate tickets",
            tickets: [["JFK", "ATL"], ["ATL", "JFK"], ["JFK", "ATL"], ["ATL", "SFO"]],
            want: ["JFK", "ATL", "JFK", "ATL", "SFO"],
        },
    ])("case: $name", ({ tickets, want }) => {
        const got = findItinerary(tickets);
        expect(got).toEqual(want);
    });
});
