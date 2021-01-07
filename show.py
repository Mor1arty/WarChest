from PIL import Image
from enum import Enum
from abc import ABC, abstractmethod


def get_pixel_coordinate(logic_pos: tuple) -> tuple:
    """
    Return pixel coordinate according to coin's logical position.Let the most left down position be (0, 0).X and y
    increase along up and right direction respectively.
    :param logic_pos: coin's logic position.
    :return: coin's pixel position according to its logic position.
    """
    init_pos = (231, 491)  # 最左下角的点的坐标
    x = init_pos[0] + logic_pos[0] * 78
    if logic_pos[0] < 4:
        y = init_pos[1] + logic_pos[0] * 44 - logic_pos[1] * 88
    else:
        y = init_pos[1] + (6 - logic_pos[0]) * 44 - logic_pos[1] * 88
    px_pos = (x, y)
    return px_pos


class GlobalCoin(Enum):
    """
    Image path of coins except unit coins.(I call them global coins)
    Default in folder imgs/.
    """
    PHOENIX_CONTROL_MARKER = "phoenix_control_marker.png"
    LION_CONTROL_MARKER = "lion_control_marker.png"
    PHOENIX_ROYAL_COIN = "phoenix_royal_coin.png"
    LION_ROYAL_COIN = "lion_royal_coin.png"


class UnitCoin(Enum):
    """
    Image path of unit coins.(Default in folder imgs/)
    """
    LIGHT_CAVALRY = "light_cavalry.png"


class Observer(ABC):
    """
    The Observer interface declares the update method, used by subjects.
    """
    @abstractmethod
    def update(self, subject: Subject) -> None:
        """
        Receive update from subject.
        :param subject:
        :return:
        """
        pass


class Subject(ABC):
    """
    The Subject interface declares a set of methods for managing subscribers.
    """
    @abstractmethod
    def attach(self, observer: Observer) -> None:
        """
        Attach an observer to the subject.
        :param observer: the observer to be added.
        :return:
        """
        pass

    @abstractmethod
    def detach(self, observer: Observer) -> None:
        """
        Detach an observer from the subject.
        :param observer: the observer to be removed.
        :return:
        """

    @abstractmethod
    def notify(self) -> None:
        """
        Notify all observers about an event.
        :return:
        """
        pass


class Coin(Subject):
    """
    The unit coin class, also called hero.
    """
    position = None
    HP = 0

    def __init__(self, faction: str, position: tuple = None):
        self.control_marker = GlobalCoin.PHOENIX_CONTROL_MARKER if faction == 'phoenix' else \
            GlobalCoin.LION_CONTROL_MARKER
        if Coin.position is None:
            Coin.position = position

    def attach(self, observer: Observer) -> None:

    def detach(self, observer: Observer) -> None:

    def notify(self) -> None:

    def placement(self, action: str):
        if Coin.position is not None:
            global ui
            if action == 'deploy':
                ui.set_coin(coin_name=self.control_marker, pos=self.position)
                Coin.on_board = 1
                self.HP = 1
            if action == 'bolster':
                pass


class UIController(Observer):
    def __init__(self):
        self.background = Image.open("imgs/background.png")
        self.set_coin(coin_name=GlobalCoin.PHOENIX_CONTROL_MARKER, pos=get_pixel_coordinate((1, 0)))
        self.set_coin(coin_name=GlobalCoin.PHOENIX_CONTROL_MARKER, pos=get_pixel_coordinate((4, 0)))
        self.set_coin(coin_name=GlobalCoin.LION_CONTROL_MARKER, pos=get_pixel_coordinate((2, 5)))
        self.set_coin(coin_name=GlobalCoin.LION_CONTROL_MARKER, pos=get_pixel_coordinate((5, 4)))

    def set_coin(self, coin_name: Enum, pos: tuple):
        img_path = f"imgs/{coin_name.value}"
        coin = UIController.__read_coin(img_path=img_path)
        self.background.paste(coin, pos, coin)

    def update(self, subject: Subject) -> None:
        self.background.show()

    @staticmethod
    def __read_coin(img_path: str) -> Image.Image:
        coin = Image.open(img_path)
        coin = coin.resize((78, 78))
        return coin


def main():
    ui = UIController()

    ui.set_coin(coin_name=UnitCoin.LIGHT_CAVALRY, pos=get_pixel_coordinate((2, 3)))
    ui.set_coin(coin_name=UnitCoin.LIGHT_CAVALRY, pos=get_pixel_coordinate((2, 3)))
    ui.set_coin(coin_name=GlobalCoin.PHOENIX_CONTROL_MARKER, pos=get_pixel_coordinate((6, 2)))

    ui.update()


if __name__ == '__main__':
    main()
