#include <bzlib.h>

// "BZ2_bzCompressInit" and "BZ2_bzCompressEnd" are called from Go
// "BZ2_bsCompress" is performed here
int bz2compress(bz_stream *s, int action, char* in, unsigned* inlen, char* out, unsigned* outlen) {
  // setting input-stream and output-stream inside the "s" stream
  s->next_in = in;
  s->avail_in = *inlen; // number of raw data [bytes] e.g. 1500
  s->next_out = out; 
  s->avail_out = *outlen; // space available for output e.g. 1000
  // compress data
  int r = BZ2_bzCompress(s, action);
  
  //// assume: compress process reduce 200 bytes => 100 bytes
  // s->avail_in: means "number-of-raw-bytes-being-left" (1300)
  // *inlen: changes meaning; now means "number-of-processed-bytes" (200)
  *inlen -= s->avail_in; 
  // s->avail_out now: space left available for output (900 bytes)
  // *outlen: changes meaning; now means "number-of-bytes-being-produced" (100)
  *outlen -= s->avail_out;
  
  return r; // status code
}