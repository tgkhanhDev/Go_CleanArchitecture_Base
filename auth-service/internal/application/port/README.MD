Why "port"? What does "port" mean?

In Onion Architecture (and Clean Architecture), "port" is short for:

Port & Adapter Architecture — also known as Hexagonal Architecture, In Onion/Hexagonal architecture, the port is:

    A contract that defines what your application expects or provides

    Divided into:
        Inbound Ports (also called Input Ports): what the outside world can call (e.g., UserService)
        Outbound Ports (also called Output Ports): what your app needs from the outside world (e.g., UserRepository, EmailSender)
It uses the idea of:

    Ports → interfaces defined by your application (abstractions): MailSender, FileStorage, EventPublisher, etc.

    Adapters → concrete implementations of those interfaces (e.g., SMTP adapter)

🧠 Think of It Like This:
Your application layer says: “I want to send an email, but I don’t care how it’s sent.”

So it defines a port (interface): EmailSender

Then, the infrastructure layer provides the adapter that knows how to actually send the email (via SMTP, AWS SES, etc.)