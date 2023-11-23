#!/bin/bash

REPO_DIR="/home/ubuntu/website_posts"
TARGET_DIR="/home/ubuntu/public"
LOG_FILE="/home/ubuntu/logs/batch.log"

touch "$LOG_FILE"
cd "$REPO_DIR"
git pull

find "$REPO_DIR" -name '*.md' | while read md_file; do
	json_file="$TARGET_DIR/$(basename "${md_file%.*}").json"
	img_file="$(basename "${md_file%.*}").png"
	title=$(basename "${md_file%.*}" | sed 's/_/ /g')
	date=$(date --iso-8601)

	if [[ -f "$json_file" ]]; then
		echo "$(date '+%Y-%m-%d %H:%M:%S') - Updating $json_file contents" >> "$LOG_FILE"
		jq --arg contents "$(<"$md_file")" '.contents = $contents' "$json_file" > tmp.json && mv tmp.json "$json_file"
	else
		echo "$(date '+%Y-%m-%d %H:%M:%S') - Creating new JSON file "$json_file"" >> "$LOG_FILE"
		jq -n \
			--arg title "$title" \
			--arg date "$date" \
			--arg img_url "$img_file" \
			--arg contents "$(<"$md_file")" \
			'{title: $title, date: $date, imgUrl: $img_url, contents: $contents}' > "$json_file"
	fi
done

echo "$(date '+%Y-%m-%d %H:%M:%S') - Copying png files to target directory" >> "$LOG_FILE"
find "$REPO_DIR" -name '*.png' -exec cp {} "$TARGET_DIR" \;

