import os
import warnings

import pandas as pd
import termcolor
import torch.utils.data
from PIL import Image
from termcolor import cprint
from torch.utils.data import Dataset
from torch.utils.data.dataset import T_co
from torchvision import datasets, transforms


class CustomDataset(Dataset):
    def __init__(self, images_dir: str, labels_df: pd.DataFrame, transform):
        cprint(labels_df["ClassName"].value_counts(), "yellow")
        self.images_dir = images_dir
        self.images_filenames = [
            f for f in os.listdir(self.images_dir) if f in labels_df.index
        ]
        self.label_files_df = labels_df.loc[self.images_filenames]
        self.transform = transform

    def __len__(self):
        return len(self.images_filenames)

    def __getitem__(self, index) -> T_co:
        filename = self.images_filenames[index]
        filepath = os.path.join(self.images_dir, filename)
        image = self.transform(Image.open(filepath))
        label = self.label_files_df.loc[filename, "ClassName"]
        return image, label


cprint("# 0. Read train and test sets", "green")

transform = transforms.Compose(
    [transforms.Resize(255), transforms.CenterCrop(224), transforms.ToTensor()]
)

train_set = CustomDataset(
    "/home/piotr/Datasets/food-classy/train_images/train_images",
    pd.read_csv("/home/piotr/Datasets/food-classy/train_img.csv", index_col=0),
    transform=transform,
)

cprint("# 1. Load datasets into pytorch dataloader", "green")

dataloader = torch.utils.data.DataLoader(
    train_set, batch_size=64, shuffle=True
)

cprint("# 2. Show examples", "green")


