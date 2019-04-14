# Concurrent Non-Blocking Cache

This is the code from my talk: **Let's build a concurrent non-blocking cache**

This tutorial will demonstrate ways to approach the design and implementation of concurrent data structures. In the process, it will clarify that there is no one size fits it all solution and different designs – such as shared variables and locks or communicating sequential processes – can make sense. It is not always obvious which approach is preferable but depending on the context, one solution can be simpler or more expressive to the specific problem domain than the other one.

The demo will use a cache as the sample project. It will fetch a handle which serves HTTP requests for computing a playlist dynamically based on existing media segments in a database. Calls to this endpoint are relatively expensive which makes it a reasonable use case for a cache.

[Video](https://youtu.be/KlDWmTcyXdA)

[Slides](https://speakerdeck.com/konradreiche/lets-build-a-concurrent-non-blocking-cache)
