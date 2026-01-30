# Lab 07 — Pipeline

## Concept

Multi-stage processing with ordered data flow.

This lab explores pipeline-based concurrency, where work flows through a sequence of processing stages. Each stage runs concurrently and communicates exclusively via message passing. Correctness depends on flow control, backpressure, and coordinated shutdown across stages.

---

## Problem Statement

Implement a multi-stage concurrent pipeline.

Input items are produced by a source, transformed by one or more intermediate stages, and consumed by a sink. Each stage operates concurrently and processes items independently, passing results downstream.

The system must ensure that all input items are processed in order, no items are lost or duplicated, and all pipeline stages terminate cleanly.

---

## Inputs

- An integer `N` representing the number of input items
- An integer `S` representing the number of pipeline stages
- A workload range controlling how long each stage takes to process an item

Inputs may be hard-coded or provided via command-line arguments.

---

## Pipeline Stage Behavior

Each stage must:

1. Begin execution
2. Receive items from the upstream stage
3. Process each item independently
4. Emit exactly one output item per input item
5. Continue until the upstream stage is exhausted
6. Terminate execution

Stages must not:
- Share mutable state with other stages
- Retain items after forwarding them

---

## Source Behavior

The source must:

1. Begin execution
2. Emit `N` input items
3. Stop emitting items
4. Terminate execution

---

## Sink Behavior

The sink must:

1. Begin execution
2. Receive all output items
3. Record or aggregate results
4. Detect completion
5. Terminate execution

---

## Output Requirements

- The program must report:
  - Total input items
  - Total output items
- Optional per-stage or per-item records may be produced
- Output ordering must match input ordering

---

## Concurrency Constraints

- Each input item must produce exactly one output item
- Items must flow through all stages in order
- No stage may drop or duplicate items
- Backpressure must propagate upstream
- The main execution context must not exit until:
  - All items reach the sink
  - All stages have terminated

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- A slow or blocked intermediate stage
- Downstream consumer failure
- Upstream burst of input items
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Pipeline clogging under slow stages
- Backpressure effects across stages
- Shutdown ordering when stages complete at different times
- Sensitivity to uneven stage workloads

---

## Success Criteria

The lab is considered complete when:

- All input items are processed in order
- All stages execute concurrently
- No goroutine leaks occur
- The pipeline drains completely
- The program terminates deterministically

---

## Notes

This lab emphasizes structured concurrency, flow control, and correctness across multiple dependent stages.

It serves as the foundation for coordination patterns and cancellation in subsequent labs.
