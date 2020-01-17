package com.xulog.algo.util;

import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.ForkJoinTask;
import java.util.concurrent.RecursiveTask;
import java.util.concurrent.atomic.AtomicInteger;

public class F {





    public static void main(String[] args) {
//        System.out.println(binomial(20, 10, 0.25));
        System.out.println(binomial(100, 50, 0.25));
        System.out.println(counter.get());
    }

    static AtomicInteger counter = new AtomicInteger(0);

    public static double binomial(int n, int k, double p) {
//        System.out.println(n + "," + k + "," + p);
        counter.getAndIncrement();
        if (n == 0 && k == 0) return 1.0;
        if (n < 0 || k < 0) return 0.0;
        return (1.0 - p) * binomial(n - 1, k, p) + p * binomial(n - 1, k - 1, p);
    }
}
