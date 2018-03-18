import (
    "fmt"
    "os"
)

func lar_booter_exit_with_catched(c *lar_gcls_inst_10___builtins_7_Catched_1_lar_intf_10___builtins_9_Throwable) {
    fmt.Fprintln(os.Stderr, lar_util_convert_lar_str_to_go_str(c.m_tb))
    os.Exit(2)
}

func lar_booter_start_prog(main_mod_init_func func (), main_func func (*[]*lar_cls_10___builtins_6_String) int32,
                           argv *[]*lar_cls_10___builtins_6_String) int {
    defer func () {
        c := lar_func_10___builtins_10_catch_base(recover())
        if c != nil {
            lar_booter_exit_with_catched(c)
        }
    }()
    lar_env_init_mod_10___builtins()
    main_mod_init_func()
    return int(main_func(argv))
}

func lar_booter_start_co(co lar_intf_10___builtins_9_Coroutine) {
    defer func () {
        c := lar_func_10___builtins_10_catch_base(recover())
        if c != nil {
            lar_booter_exit_with_catched(c)
        }
    }()
    co.method_run()
}
