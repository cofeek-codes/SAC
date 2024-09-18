import FreeSimpleGUI as sg


def center_cursor():
    pass


def config():
    sg.user_settings_filename(filename="./config.json")
    layout = [
        [
            sg.Text("Start/Stop Button"),
            sg.Input(sg.user_settings_get_entry("-start_stop_button-", "a")),
        ],
        [
            sg.Text("Exit Button"),
            sg.Input(sg.user_settings_get_entry("-exit_button-", "b")),
        ],
        [
            sg.Text("Delay"),
            sg.Input(sg.user_settings_get_entry("-delay-", float(0.5))),
        ],
        [sg.B("Launch"), sg.B("Exit")],
    ]

    window = sg.Window("SAC Config", layout)

    while True:
        event, values = window.read()

        if event == sg.WIN_CLOSED or event == "Exit":
            # @TODO: shutdown everything here
            break

        elif event == "Launch":
            print("saving and launching")

        window.close()


def main():
    config()


if __name__ == "__main__":
    main()
