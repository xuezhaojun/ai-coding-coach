import { describe, it, expect } from 'vitest';
import { canAttendMeetings } from './meetingRooms';

describe('MeetingRooms', () => {
    it('overlapping', () => {
        expect(canAttendMeetings([[0, 30], [5, 10], [15, 20]])).toBe(false);
    });

    it('no overlap', () => {
        expect(canAttendMeetings([[7, 10], [2, 4]])).toBe(true);
    });

    it('empty', () => {
        expect(canAttendMeetings([])).toBe(true);
    });

    it('single meeting', () => {
        expect(canAttendMeetings([[1, 5]])).toBe(true);
    });

    it('touching boundaries', () => {
        expect(canAttendMeetings([[1, 5], [5, 10]])).toBe(true);
    });

    it('same time', () => {
        expect(canAttendMeetings([[1, 5], [1, 5]])).toBe(false);
    });

    it('sequential', () => {
        expect(canAttendMeetings([[0, 1], [1, 2], [2, 3]])).toBe(true);
    });
});
