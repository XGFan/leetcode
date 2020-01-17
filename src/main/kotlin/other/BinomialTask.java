package com.xulog.algo.util;

import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.RecursiveTask;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.AtomicLong;

public class BinomialTask extends RecursiveTask<Double> {

    static AtomicLong counter = new AtomicLong();

    private static final long serialVersionUID = 6253771003381008573L;


    private int n;
    private int k;
    private double p;


    public BinomialTask(int n, int k, double p) {
        counter.getAndIncrement();
        this.n = n;
        this.k = k;
        this.p = p;
    }


    @Override
    protected Double compute() {
        if (n == 0 && k == 0) return 1.0;
        else if (n < 0 || k < 0) return 0.0;
        else {
            if (1.0 - p != 0) {
                BinomialTask leftTask = new BinomialTask(n - 1, k, p);
                BinomialTask rightTask = new BinomialTask(n - 1, k - 1, p);
                leftTask.fork();
                rightTask.fork();
                return (1.0 - p) * leftTask.join() + p * rightTask.join();
            } else {
                BinomialTask rightTask = new BinomialTask(n - 1, k - 1, p);
                rightTask.fork();
                return p * rightTask.join();
            }
        }
    }


    /**
     * 20 10 0.25
     * 0.009922275279677706
     * 2433071
     */
    public static void main(String[] args) {
        ForkJoinPool pool = new ForkJoinPool();
        final BinomialTask binomialTask = new BinomialTask(100, 50, 0.25);
        final Double invoke = pool.invoke(binomialTask);
        System.out.println(invoke);
        System.out.println(counter);
    }
}
