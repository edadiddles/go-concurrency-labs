# Lab 16 — Throughput vs Latency

## Concept

Performance tradeoffs under concurrency.

This lab explores the inherent tension between maximizing throughput and minimizing latency in concurrent systems. By varying buffering, batching, and scheduling strategies, the system exhibits different performance characteristics under load.

---

## Problem Statement

Implement a concurrent processing system that accepts work items, processes them, and emits results. The system must allow configuration of concurrency limits, buffering, and scheduling behavior.

The goal is to observe how design choices affect end-to-end latency and total throughput under different workload patterns.

---

## Inputs

- Concurrency limit
- Input buffer size
- Workload size and submission rate
- Task execution cost distribution
- Output sink behavior

Inputs may be hard-coded or provided via command-line arguments.

---

## Work Item Behavior

Each work item must:

1. Record the time it enters the system
2. Perform a unit of work (blocking or CPU-bound)
3. Record the time it exits the system
4. Emit a completion record containing latency metrics

Work items must not:
- Depend on execution order
- Access shared mutable state without synchronization
- Assume immediate processing upon submission

---

## System Behavior

The system must:

- Process multiple work items concurrently
- Support configurable buffering between stages
- Apply backpressure when overloaded
- Remain stable under sustained load

---

## Output Requirements

The program must produce metrics including:

- Per-item latency
- Aggregate throughput
- Queue depth over time (optional)
- Percentile latency statistics (optional)

Output must be readable and consistent across runs.

---

## Concurrency Constraints

- Concurrency must remain within configured limits
- No unbounded resource growth
- No work item may be processed more than once
- No work item may be silently dropped unless explicitly configured

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Sudden workload spikes
- Slow or blocking downstream processing
- Uneven task execution times
- Buffer size changes during runtime (if supported)

---

## Observations to Make

While running the program, observe:

- Latency growth under increasing load
- Throughput saturation points
- Queue buildup behavior
- Tail latency sensitivity

---

## Success Criteria

The lab is considered complete when:

- Throughput and latency metrics are measurable and reproducible
- Backpressure behavior is observable
- Tradeoffs between buffering and responsiveness are evident
- System remains stable under stress
- No goroutine leaks or deadlocks occur

---

## Notes

This lab synthesizes all prior concurrency concepts: worker pools, queues, cancellation, and contention. It emphasizes measurement, not optimization, and encourages experimentation with system design tradeoffs.
