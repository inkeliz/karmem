#include "stdint.h"
#include "stdlib.h"
#include "string.h"
#include "stdbool.h"

typedef struct {
    uint8_t *memory;
    uint32_t capacity;
    uint32_t length;
    uint8_t is_fixed;
} KarmemWriter;

KarmemWriter KarmemNewWriter(size_t capacity) {
    return (KarmemWriter) {.memory = (uint8_t *)malloc(capacity), .capacity = (uint32_t) capacity, .length = 0, .is_fixed = 0};
}

KarmemWriter KarmemNewFixedWriter(uint8_t * mem, size_t capacity) {
    return (KarmemWriter) {.memory = mem, .capacity = capacity, .length = 0, .is_fixed = 1};
}

uint32_t KarmemWriterAlloc(KarmemWriter * w, size_t size) {
    uint32_t ptr = w->length;
    uint32_t total = ptr + (uint32_t) size;
    if (total > w->capacity) {
        if (w->is_fixed > 0) {
            return 0xFFFFFFFF;
        }
        uint32_t capacity_target = w->capacity * 2;
        if (total > capacity_target) {
            capacity_target = total;
        }
        w->memory = realloc(w->memory, total);
        w->capacity = capacity_target;
    }
    w->length = total;

    return ptr;
}

void KarmemWriterWriteAt(KarmemWriter * w, uint32_t offset, const void * src, size_t length) {
    memcpy(&w->memory[offset], src, length);
}

void KarmemWriterReset(KarmemWriter * w) {
    w->length = 0;
}

void KarmemWriterFree(KarmemWriter * w) {
    free(w->memory);
}

typedef struct {
    uint8_t *pointer;
    uint32_t length;
    uint32_t capacity;
} KarmemReader;

KarmemReader KarmemNewReader(uint8_t * memory, uint32_t length) {
    return (KarmemReader) {.pointer = memory, .length = length, .capacity = length};
}

bool KarmemReaderIsValidOffset(KarmemReader * r, uint32_t offset, uint32_t size) {
    return r->length >= offset + size;
}

// SetSize re-defines the bounds of the slice, useful when the
// backend slice is being re-used for multiples contents.
uint8_t KarmemReaderSetSize(KarmemReader * r, uint32_t size) {
    if (size > r->capacity) {
        return 0;
    }
    r->length = size;
    return 1;
}
