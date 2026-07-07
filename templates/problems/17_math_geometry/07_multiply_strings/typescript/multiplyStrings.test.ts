import { describe, it, expect } from 'vitest';
import { multiply } from './multiplyStrings';

describe('MultiplyStrings', () => {
    it('example 1', () => {
        expect(multiply("2", "3")).toBe("6");
    });

    it('example 2', () => {
        expect(multiply("123", "456")).toBe("56088");
    });

    it('multiply by zero', () => {
        expect(multiply("0", "52")).toBe("0");
    });

    it('both zeros', () => {
        expect(multiply("0", "0")).toBe("0");
    });

    it('single digits', () => {
        expect(multiply("9", "9")).toBe("81");
    });

    it('large numbers', () => {
        expect(multiply("999", "999")).toBe("998001");
    });

    it('one and number', () => {
        expect(multiply("1", "12345")).toBe("12345");
    });

    it('overflow case', () => {
        expect(multiply("498828660196", "840477629533")).toBe("419254329864656431168468");
    });
});
