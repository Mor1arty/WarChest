from PIL import Image


class UIController(object):
    def __init__(self):
        self.background = Image.open("imgs/background.png")

    def set_coin(self, coin_path: str, pos: tuple):
        coin = UIController.__read_coin(img_name=coin_path)
        self.background.paste(coin, pos, coin)

    def show(self):
        self.background.show()

    @staticmethod
    def __read_coin(img_name: str) -> Image.Image:
        coin = Image.open(f"imgs/{img_name}")
        coin = coin.resize((78, 78))
        return coin


def main():
    UI = UIController()
    UI.set_coin(coin_path="ppoint-removebg-preview.png", pos=(309, 534))
    UI.set_coin(coin_path="ppoint-removebg-preview.png", pos=(543, 578))

    UI.show()


if __name__ == '__main__':
    main()
