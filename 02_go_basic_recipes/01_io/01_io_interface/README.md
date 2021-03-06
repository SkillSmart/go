The Copy() function copies bytes between interfaces and treats that data like a stream. Thinking of data as streams has 
a lot of practical uses, especially when working with network traffic or filesystems. The Copy() function also creates 
a MultiWriter interface that combines two writer streams and writes to them twice using ReadSeeker. If a Reader 
interface was used instead, rather than seeing exampleexample, you would only see example despite copying to the 
MultiWriter interface twice. 

You can also use a buffered write if your stream is not fitted into the memory.The 
PipeReader and PipeWriter structures implement the io.Reader and io.Writer interfaces. They're connected, creating an 
in-memory pipe. The primary purpose of a pipe is to read from a stream while simultaneously writing from the same stream 
to a different source. In essence, it combines the two streams into a pipe.Go interfaces are a clean abstraction to wrap 
data that performs common operations. This is made apparent when doing I/O operations, and so the io package is a great 
resource for ing about interface composition. The pipe package is often underused, but provides great flexibility 
with thread safety when linking input and output streams.