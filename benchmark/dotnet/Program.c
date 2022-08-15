#include <mono-wasi/driver.h>
#include <assert.h>
#include <string.h>
#include <stdint.h>
#include <stdio.h>

// Class (similar to Java's JNI):
// Initialize when Ready is called.
MonoObject* _CLASS_Benchmark = NULL;

// Method IDS (similar to Java's JNI):
// Initialize when Ready is called.
MonoMethod* _ID_InputMemoryPointer;
MonoMethod* _ID_OutputMemoryPointer;
MonoMethod* _ID_KBenchmarkEncodeObjectAPI;
MonoMethod* _ID_KBenchmarkDecodeObjectAPI;
MonoMethod* _ID_KBenchmarkDecodeSumVec3;

__attribute__((export_name("InputMemoryPointer")))
uint32_t InputMemoryPointer() {
    MonoObject* err;
    MonoObject* res = mono_wasm_invoke_method (_ID_InputMemoryPointer, _CLASS_Benchmark, NULL, &err);
    if (err != NULL) {
        assert(err);
    }
    return *(uint32_t*)mono_object_unbox(res);
}

__attribute__((export_name("OutputMemoryPointer")))
uint32_t OutputMemoryPointer() {
    MonoObject* err;
    MonoObject* res = mono_wasm_invoke_method (_ID_OutputMemoryPointer, _CLASS_Benchmark, NULL, &err);
    if (err != NULL) {
        assert(err);
    }
    return *(uint32_t*)mono_object_unbox(res);
}

__attribute__((export_name("KBenchmarkEncodeObjectAPI")))
void KBenchmarkEncodeObjectAPI() {
    MonoObject* err;
    MonoObject* res = mono_wasm_invoke_method (_ID_KBenchmarkEncodeObjectAPI, _CLASS_Benchmark, NULL, &err);
    if (err != NULL) {
        assert(err);
    }
}

__attribute__((export_name("KBenchmarkDecodeObjectAPI")))
void KBenchmarkDecodeObjectAPI(uint32_t size) {
    void* params[] = { &size };
    MonoObject* err;
    MonoObject* res = mono_wasm_invoke_method (_ID_KBenchmarkDecodeObjectAPI, _CLASS_Benchmark, params, &err);
    if (err != NULL) {
        assert(err);
    }
}

__attribute__((export_name("KBenchmarkDecodeSumVec3")))
float KBenchmarkDecodeSumVec3(uint32_t size) {
    void* params[] = { &size };
    MonoObject* err;
    MonoObject* res = mono_wasm_invoke_method (_ID_KBenchmarkDecodeSumVec3, _CLASS_Benchmark, params, &err);
    if (err != NULL)
    {
        assert(err);
    }
    return *(float*)mono_object_unbox(res);
}


void ready_benchmark(MonoObject* cls) {
    _CLASS_Benchmark = cls;

    _ID_InputMemoryPointer = lookup_dotnet_method("dotnet.dll", "dotnet", "Benchmark", "InputMemoryPointer", -1);
    _ID_OutputMemoryPointer = lookup_dotnet_method("dotnet.dll", "dotnet", "Benchmark", "OutputMemoryPointer", -1);
    _ID_KBenchmarkEncodeObjectAPI = lookup_dotnet_method("dotnet.dll", "dotnet", "Benchmark", "KBenchmarkEncodeObjectAPI", -1);
    _ID_KBenchmarkDecodeObjectAPI = lookup_dotnet_method("dotnet.dll", "dotnet", "Benchmark", "KBenchmarkDecodeObjectAPI", -1);
    _ID_KBenchmarkDecodeSumVec3 = lookup_dotnet_method("dotnet.dll", "dotnet", "Benchmark", "KBenchmarkDecodeSumVec3", -1);
}

void dotnet_benchmark_init() {
    mono_add_internal_call ("dotnet.Benchmark::Ready", ready_benchmark);
}