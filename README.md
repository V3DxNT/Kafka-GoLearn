# Kafka Producer-Consumer Demo

A small hands-on project demonstrating the basic **Producer-Consumer pattern using Apache Kafka**.

The goal of this project is to show how a producer can publish messages to a Kafka topic and how a consumer can subscribe to that topic and receive those messages.

It is intentionally kept simple so that the core Kafka concepts are easy to understand without adding unnecessary complexity.

## 📌 Project Overview

This project contains two main components:

* **Producer** — sends messages to a Kafka topic.
* **Consumer** — listens to the same topic and receives the messages.

The basic flow looks like this:

```text
        ┌──────────┐
        │ Producer │
        └────┬─────┘
             │
             │ Message
             ▼
      ┌──────────────┐
      │ Apache Kafka │
      │    Topic     │
      └──────┬───────┘
             │
             │ Message
             ▼
        ┌──────────┐
        │ Consumer │
        └──────────┘
```

For example, when the producer sends:

```text
Hello from Kafka!
```

Kafka stores the message in the configured topic, and the consumer receives:

```text
Received: Hello from Kafka!
```

## 🎯 Purpose

This mini project is intended to demonstrate the fundamentals of Kafka, including:

* Creating a Kafka producer
* Creating a Kafka consumer
* Sending messages to a Kafka topic
* Receiving messages from a Kafka topic
* Understanding the basic producer → Kafka → consumer workflow
* Getting familiar with Kafka topics and consumer groups

## 🛠️ Technologies Used

* **Apache Kafka**
* **[Add your programming language/framework here]**
* Kafka Producer API
* Kafka Consumer API

> Replace the programming language/framework above with the technology used in your implementation, such as Java, Spring Boot, Python, Node.js, etc.

## 📂 Project Structure

A typical structure for this project can look like:

```text
kafka-producer-consumer/
│
├── producer/
│   └── Producer application
│
├── consumer/
│   └── Consumer application
│
├── README.md
└── ...
```

The `producer` is responsible for publishing messages, while the `consumer` continuously listens for new messages.

## ⚙️ How It Works

### 1. Kafka Starts

First, a Kafka broker needs to be running locally or on a server.

The producer and consumer both connect to the Kafka broker.

### 2. Producer Sends a Message

The producer creates a message and sends it to a Kafka topic.

For example:

```text
Topic: demo-topic
Message: Hello Kafka!
```

### 3. Kafka Stores the Message

Kafka receives the message and stores it in the specified topic.

The message remains available according to the topic's retention configuration.

### 4. Consumer Receives the Message

The consumer subscribes to `demo-topic`.

When a message is available, the consumer reads and processes it.

Example output:

```text
Consumer started...
Received message: Hello Kafka!
```

## 🚀 Getting Started

### Prerequisites

Before running the project, make sure you have:

* Apache Kafka installed and configured
* A running Kafka broker
* The required runtime for the chosen programming language
* Kafka's required dependencies installed for the project

### 1. Start Kafka

Start your Kafka broker using your local Kafka setup.

Make sure the broker is available at the address configured by your application, for example:

```text
localhost:9092
```

### 2. Create the Topic

Create a topic for the demo:

```text
demo-topic
```

The exact command depends on your Kafka version and setup.

### 3. Start the Consumer

Start the consumer application first.

You should see something similar to:

```text
Consumer started...
Waiting for messages...
```

### 4. Start the Producer

Run the producer application and send a message:

```text
Hello from Kafka!
```

The consumer should then receive it:

```text
Received: Hello from Kafka!
```

## 🧪 Example

### Producer

```text
Sending message: Hello Kafka!
Sending message: This is my first Kafka message.
Sending message: Producer -> Kafka -> Consumer
```

### Consumer

```text
Received message: Hello Kafka!
Received message: This is my first Kafka message.
Received message: Producer -> Kafka -> Consumer
```

This demonstrates the complete message flow:

```text
Producer
   │
   │  "Hello Kafka!"
   ▼
Kafka Topic
   │
   │  "Hello Kafka!"
   ▼
Consumer
```

## 🔑 Key Kafka Concepts

### Producer

A **producer** is an application that publishes messages to Kafka.

In this project, the producer is responsible for sending messages to `demo-topic`.

### Consumer

A **consumer** is an application that reads messages from Kafka.

The consumer subscribes to the topic and processes messages as they become available.

### Topic

A **topic** is a logical category where Kafka stores messages.

This project uses:

```text
demo-topic
```

### Consumer Group

Consumers can belong to a **consumer group**. Kafka uses consumer groups to distribute messages between multiple consumers.

For this basic project, a single consumer is enough to demonstrate the concept.

## 📋 Example Configuration

A typical local setup might use:

```text
Kafka Broker: localhost:9092
Topic:        demo-topic
Consumer Group: demo-consumer-group
```

These values can be changed according to your environment.

## 📚 What This Project Demonstrates

After running this project, you should have a basic understanding of:

1. How a Kafka producer publishes messages.
2. How Kafka topics are used to organize messages.
3. How a consumer subscribes to a topic.
4. How messages move from a producer to a consumer through Kafka.
5. The basic idea behind Kafka consumer groups.

## 🔮 Possible Improvements

This project is intentionally minimal, but it can be extended with features such as:

* Sending JSON messages
* Multiple producers
* Multiple consumers
* Multiple partitions
* Consumer groups with multiple consumers
* Message keys
* Error handling and retries
* Docker-based Kafka setup
* REST API for producing messages
* Logging and monitoring

## 🤝 Contributing

This is primarily a learning project, but feel free to experiment with it and extend it with additional Kafka features.

## 📄 License

This project is intended for educational and demonstration purposes.
