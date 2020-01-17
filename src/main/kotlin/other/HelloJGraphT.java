//package com.xulog.algo.util;
//
//import org.jgrapht.alg.ConnectivityInspector;
//import org.jgrapht.alg.KosarajuStrongConnectivityInspector;
//import org.jgrapht.alg.interfaces.StrongConnectivityAlgorithm;
//import org.jgrapht.graph.DefaultDirectedGraph;
//import org.jgrapht.graph.DefaultEdge;
//
//import java.util.List;
//import java.util.Set;
//
//public class HelloJGraphT {
//    private DefaultDirectedGraph<String, DefaultEdge> directedGraph;
//
//    public HelloJGraphT() {
//        directedGraph = new DefaultDirectedGraph<String, DefaultEdge>(DefaultEdge.class);
//        directedGraph.addVertex("a");
//        directedGraph.addVertex("b");
//        directedGraph.addVertex("c");
//        directedGraph.addVertex("d");
//        directedGraph.addVertex("e");
//        directedGraph.addVertex("f");
//        directedGraph.addVertex("h");
//        directedGraph.addVertex("i");
//        directedGraph.addEdge("a", "b");
//        directedGraph.addEdge("b", "c");
//        directedGraph.addEdge("c", "d");
//        directedGraph.addEdge("d", "a");
//        directedGraph.addEdge("c", "e");
//        directedGraph.addEdge("f", "h");
//        directedGraph.addEdge("f", "i");
//    }
//
//    public void testStrongConn() {
//        StrongConnectivityAlgorithm<String, DefaultEdge> scAlg =
//                new KosarajuStrongConnectivityInspector<String, DefaultEdge>(directedGraph);
//        List<Set<String>> stronglyConnetedSet =
//                scAlg.stronglyConnectedSets();
//
//        System.out.println("Strongly connected components:");
//        for (int i = 0; i < stronglyConnetedSet.size(); i++) {
//            System.out.println(stronglyConnetedSet.get(i));
//        }
//        System.out.println();
//
//    }
//
//    public void testWeakConn() {
//        ConnectivityInspector<String, DefaultEdge> connectivityInspector = new ConnectivityInspector<String, DefaultEdge>(directedGraph);
//        List<Set<String>> weaklyConnectedSet = connectivityInspector.connectedSets();
//        System.out.println("Weakly connected components:");
//        for (int i = 0; i < weaklyConnectedSet.size(); i++) {
//            System.out.println(weaklyConnectedSet.get(i));
//        }
//        System.out.println();
//
//    }
//
//    public static void main(String args[]) {
//        final HelloJGraphT test = new HelloJGraphT();
//        test.testStrongConn();
//        test.testWeakConn();
//    }
//}