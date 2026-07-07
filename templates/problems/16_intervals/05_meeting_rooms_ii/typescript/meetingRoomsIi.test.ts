import { describe, it, expect } from 'vitest';
import { minMeetingRooms } from './meetingRoomsIi';

describe('MeetingRoomsIi', () => {
    it('two rooms', () => {
        expect(minMeetingRooms([[0, 30], [5, 10], [15, 20]])).toBe(2);
    });

    it('one room', () => {
        expect(minMeetingRooms([[7, 10], [2, 4]])).toBe(1);
    });

    it('all overlap', () => {
        expect(minMeetingRooms([[1, 5], [2, 6], [3, 7]])).toBe(3);
    });

    it('single meeting', () => {
        expect(minMeetingRooms([[1, 10]])).toBe(1);
    });

    it('sequential', () => {
        expect(minMeetingRooms([[0, 5], [5, 10], [10, 15]])).toBe(1);
    });

    it('nested', () => {
        expect(minMeetingRooms([[1, 10], [2, 5], [6, 9]])).toBe(2);
    });

    it('empty', () => {
        expect(minMeetingRooms([])).toBe(0);
    });
});
