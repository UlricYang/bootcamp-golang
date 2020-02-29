import os
import re

import pandas as pd
from beeprint import pp as ppt

from zhtools.langconv import Converter


class Organization(object):
    """docstring for Organization."""

    def __init__(self, foldername="files"):
        super(Organization, self).__init__()
        self.foldername = foldername
        self.txtfiles = sorted(
            [
                os.path.join(root, f)
                for root, dirs, files in os.walk(
                    os.path.join(os.path.dirname(os.path.abspath(__file__)), self.foldername)
                )
                for f in files
            ]
        )

    def process_line(self, line, filepath):
        text_cht, text_chs, link, pswd = str(), str(), str(), str()
        try:
            text_cht = re.findall(r"(.*)https", line).pop().strip()
            text_chs = Converter("zh-hans").convert(text_cht)
            link = re.findall(r"https://pan.baidu.com/s/[\w\-]+", line).pop()
            pswd = re.findall(r"提取码：(\w{4,4})", line).pop() if "提取码" in line else str()
        except Exception as e:
            ppt(e)
            ppt(filepath)
            ppt(line)
        return text_cht, text_chs, link, pswd

    def extract_single(self, filepath):
        tag = "-".join(re.findall(r"\d+", os.path.basename(filepath)))
        with open(filepath, "r") as fp:
            lines = fp.readlines()
            lines = [line.strip("\n") for line in filter(lambda line: line != "\n", lines)]
            lines = list(filter(lambda l: "提取码" not in l[0], zip(lines[:-1], lines[1:])))
            lines = list(map(lambda l: " ".join(l) if "提取码" in l[1] else l[0], lines))
        day = str()
        results = list()
        for line in lines:
            if re.match(r"(\d+)\.(\d+)\.(\d+)", line):
                day = "-".join(re.search(r"(\d+)\.(\d+)\.(\d+)", line).groups())
            else:
                results.append((day, *self.process_line(line, filepath)))
        return {tag: results}

    def extract_multiple(self):
        r = list(map(lambda f: self.extract_single(f), self.txtfiles))
        res = dict()
        for item in r:
            res.update(item)
        return res

    def write_xlsx(self, res, filename):
        writer = pd.ExcelWriter(filename)
        for k in sorted(res.keys()):
            v = res[k]
            df = pd.DataFrame(
                {
                    "日期": [each[0] for each in v],
                    "描述繁体": [each[1] for each in v],
                    "描述简体": [each[2] for each in v],
                    "链接": [each[3] for each in v],
                    "提取码": [each[4] for each in v],
                }
            )
            df.to_excel(writer, sheet_name=k)
        writer.save()

    def organize(self, filename):
        res = self.extract_multiple()
        self.write_xlsx(res, filename)


def main():
    obj = Organization()
    obj.organize("汇总2.xlsx")


if __name__ == "__main__":
    main()
