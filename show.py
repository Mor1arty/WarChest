from PIL import Image


def read_coin(img_name: str) -> Image.Image:
    coin = Image.open(f"imgs/{img_name}")
    coin = coin.resize((78, 78))
    return coin


def set_coin(coin: Image.Image, pos: tuple, bg: Image.Image):
    alpha = 1.0
    new_coin = Image.new("RGBA", coin.size)
    new_coin = Image.blend(new_coin, coin, alpha)
    bg.paste(new_coin, pos, new_coin)


def main():
    coin = read_coin("ppoint-removebg-preview.png")
    bg = Image.open("imgs/background.png")
    set_coin(coin=coin, pos=(307, 534), bg=bg)
    set_coin(coin=coin, pos=(543, 578), bg=bg)
    bg.show()


if __name__ == '__main__':
    main()