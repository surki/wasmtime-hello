from wasmer import wasi, Store, Module, Instance

store = Store()
module = Module(store, open('./hello-world.wasm', 'rb').read())

wasi_version = wasi.get_version(module, strict=True)
wasi_env = wasi.StateBuilder('hello-world').finalize()

import_obj = wasi_env.generate_import_object(store, wasi_version)
instance = Instance(module, import_obj)
instance.exports.hello_world()