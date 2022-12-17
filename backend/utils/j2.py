import os

import argparse
import yaml
import jinja2


def env_override(value, key):
    return os.getenv(key, value)


def load_filters(env):
    env.filters["env_override"] = env_override
    return env


def load_data_files(files):
    data = {}
    for fin in files:
        data.update(yaml.safe_load(fin))
    return data


def main():
    parser = argparse.ArgumentParser(
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    parser.add_argument(
        "-t",
        "--template-path",
        type=str,
        default=os.getcwd(),
        help="path to template directory",
    )
    parser.add_argument(
        "template_name",
        type=str,
        help="jinja2 template name",
    )
    parser.add_argument(
        "data_file",
        nargs="+",
        type=argparse.FileType("r"),
        help="YAML-formatted data file",
    )

    args = parser.parse_args()

    j2_env = load_filters(
        jinja2.Environment(
            autoescape=True, loader=jinja2.FileSystemLoader(args.template_path)
        )
    )

    temp = j2_env.get_template(args.template_name)

    print(temp.render(load_data_files(args.data_file)))


if __name__ == "__main__":
    main()
