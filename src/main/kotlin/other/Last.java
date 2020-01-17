package com.xulog.algo.util;

public class Last {


    public static Long f(Long m, Long n) {
        if (m == 0L && n == 0L) {
            return 1L;
        }
        if (m < 0L || n < 0L) {
            return 1L;
        }
        if (n == 0) {
            return 3L;
        }
        Long total = 0L;
        for (long i = 0; i < n; i++) {
            total += f(m - 1, i);
        }
        return n + f(m, 0L) + total;

    }

    public static void main(String[] args) {

        System.out.println(f(10L, 20L));
        System.out.println(f(50L, 100L));

    }
}
